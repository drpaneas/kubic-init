#!/bin/bash

# NOTE: these variables will be replaced by Terraform
IMAGE="${kubic_init_image_name}"
IMAGE_FILENAME="${kubic_init_image_tgz}"
RUNNER=${kubic_init_runner}
EXTRA_ARGS="${kubic_init_extra_args}"
EXTRA_RPMS="${extra_rpms}"
EXTRA_RPMS_REPO="${extra_rpms_repo}"

###################################################################################

log() { echo ">>> $@" ; }

set_var() {
    var="$1"
    value="$2"
    file="$3"

    if [ -f $file ] ; then
        log "Setting $var=$value in $file..."
        sed -i 's|^\('"$var"'\)\ *=.*|\1 = '"$value"'|' $file
    else
        log "Creating new file $file, with $var=$value..."
        echo "$var = $value" > $file
    fi
    # TODO: add a case for where the file exists but the var doesn't
}

print_file() {
    log "Contents of $1:"
    log "------------------------------------"
    cat $1 | awk '{ print ">>> " $0 }'
    log "------------------------------------"
}

# make sure a repo is installed
ensure_repo() {
    repo_url=$1
    repo_name="kubic-extra"

    log "Ensuring repository $repo_url is installed"
    zypper lr $repo_name || zypper ar --refresh --no-gpgcheck --enable $repo_url $repo_name
}

# make sure a package is installed
# if it is not,  try to install it from the given repo
ensure_rpms() {
    log "Ensuring $@ are installed"
    rpm -q $rpm || zypper in -y --no-recommends --force-resolution $@
}

ensure_dns() {
    if host www.opensuse.org >/dev/null ; then
        log "DNS seems to work fine"
    else
        log "WARNING: setting up DNS with resolver 8.8.8.8"
        rm -f /etc/resolv.conf
        echo "nameserver 8.8.8.8" > /etc/resolv.conf
    fi
}

###################################################################################

log "Ensuring DNS is working"
ensure_dns

log "Making sure kubic-init is not running..."
systemctl stop kubic-init  >/dev/null 2>&1 || /bin/true

[ -z "$EXTRA_RPMS_REPO" ] || ensure_repo $EXTRA_RPMS_REPO
[ -z "$EXTRA_RPMS"      ] || ensure_rpms $EXTRA_RPMS

log "Setting runner as $RUNNER..."
[ -x /usr/bin/$RUNNER ] || ( log "FATAL: /usr/bin/$RUNNER does not exist!!!" ; exit 1 ; )
set_var KUBIC_INIT_RUNNER /usr/bin/$RUNNER /etc/sysconfig/kubic-init

log "Removing any previous kubic-init image..."
/usr/bin/$RUNNER rmi kubic-init  >/dev/null 2>&1 || /bin/true

log "Setting up network..."
echo br_netfilter > /etc/modules-load.d/br_netfilter.conf
modprobe br_netfilter
sysctl -w net.ipv4.ip_forward=1

log "Setting up crio..."
mkdir -p /etc/crio
set_var plugin_dir \"/var/lib/kubelet/cni/bin\" /etc/crio/crio.conf
echo 'runtime-endpoint: unix:///var/run/crio/crio.sock' > /etc/crictl.yaml

log "Setting up storage..."
set_var driver \"btrfs\" /etc/containers/storage.conf

[ "$RUNNER" = "podman" ] && IMAGE="localhost/$IMAGE"
log "Setting image as $IMAGE"
set_var KUBIC_INIT_IMAGE "\"$IMAGE\"" /etc/sysconfig/kubic-init

if [ -n "$EXTRA_ARGS" ] ; then
    log "Setting kubic-init extra args = $EXTRA_ARGS"
    set_var KUBIC_INIT_EXTRA_ARGS "\"$EXTRA_ARGS\"" /etc/sysconfig/kubic-init
fi

print_file /etc/sysconfig/kubic-init
print_file /etc/kubic/kubic-init.yaml

[ -d /var/lib/etcd ] || ( log "Creating etcd staorage..." ;  mkdir -p /var/lib/etcd ; )

log "Loading the kubic-init image with $RUNNER from /tmp/$IMAGE_FILENAME..."
while ! /usr/bin/$RUNNER load -i /tmp/$IMAGE_FILENAME ; do
    log "(will try to load the kubic-init image again)"
    sleep 5
done

log "Enabling the kubic-init service..."
sysctl --system
systemctl daemon-reload
systemctl enable kubelet

[ "$RUNNER" = "podman" ] && \
    systemctl enable --now crio || \
    systemctl enable --now docker

systemctl enable --now kubic-init
