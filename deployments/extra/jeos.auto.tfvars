#
# extra vars for using JeOS instead of the MicroOS/kubic images
#
# usage: include these vars before running one of 
#        the `make tf-*` targets by doing, for example:
#
#        (cd deployments/tf-libvirt-full && ln -s ../extra/jeos.auto.tfvars)
#

# img_url_base = "https://download.opensuse.org/distribution/openSUSE-stable/jeos/"
img_url_base = "https://download.opensuse.org/repositories/Cloud:/Images:/Leap_15.0/images/"

# a regex that should match one (and only one) image in the <img_url_base>
# img_regex = "JeOS.*kvm-and-xen-Current"
img_regex = "openSUSE-Leap-15.0-OpenStack"

# some extra RPMs to install
extra_rpms = "podman kubernetes-kubelet cri-o cri-tools"

# we must use the Tumbleweed repo: it has all the necessary packages
extra_rpms_repo = "https://download.opensuse.org/repositories/devel:/kubic/openSUSE_Tumbleweed/"
