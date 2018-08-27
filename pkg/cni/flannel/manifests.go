package flannel

const (
	FlannelConfigMap19 = `
kind: ConfigMap
apiVersion: v1
metadata:
  name: flannel-plugin-config-map
  namespace: kube-system
  labels:
    tier: node
    app: flannel
data:
  cni-conf.json: |
    {
      "name":"cbr0",
      "cniVersion":"0.3.1",
      "plugins":[
        {
          "type":"flannel",
          "delegate":{
              "forceAddress":true,
              "isDefaultGateway":true
          }
        },
        {
          "type":"portmap",
          "capabilities":{
            "portMappings":true
          }
        }
      ]
    }
  net-conf.json: |
    {
      "Network":   "{{ .Network }}",
      "SubnetLen": {{ .SubnetLen }},
      "Backend":
      {
        "Type": "{{ .Backend }}"
      }
    }
`

	FlannelDaemonSet19 = `
apiVersion: apps/v1beta2
kind: DaemonSet
metadata:
  name: kube-flannel
  namespace: kube-system
  labels:
    tier: node
    k8s-app: flannel
spec:
  selector:
    matchLabels:
      tier: node
      k8s-app: flannel
  template:
    metadata:
      labels:
        tier: node
        k8s-app: flannel
    spec:
      serviceAccountName: flannel
      initContainers:
      - name: install-cni-conf
        image: {{ .Image }}
        command:
          - /bin/sh
          - "-c"
          - "cp -f /etc/kube-flannel/cni-conf.json /host/etc/cni/net.d/10-flannel.conflist"
        volumeMounts:
        - name: flannel-plugin-config
          mountPath: /etc/kube-flannel/
        - name: host-cni-conf
          mountPath: /host/etc/cni/net.d
      - name: install-cni-bin
        image: {{ .Image }}
        command:
          - /bin/sh
          - "-c"
          - "cp -f /usr/lib/cni/* /host/opt/cni/bin/"
        volumeMounts:
        - name: host-cni-bin
          mountPath: /host/opt/cni/bin/
      containers:
      - name: kube-flannel
        image: {{ .Image }}
        command:
          - /usr/sbin/flanneld
          - "--ip-masq"
          - "--kube-subnet-mgr"
          - "--v={{ .LogLevel }}"
          - "--iface=$(POD_IP)"
          - "--healthz-ip=$(POD_IP)"
          - "--healthz-port={{ .HealthzPort }}"
        securityContext:
          privileged: true
        ports:
        - name: healthz
          containerPort: {{ .HealthzPort }}
        readinessProbe:
          httpGet:
            path: /healthz
            port: healthz
        livenessProbe:
          initialDelaySeconds: 10
          timeoutSeconds: 10
          httpGet:
            path: /healthz
            port: healthz
        env:
        - name: POD_IP
          valueFrom:
            fieldRef:
              fieldPath: status.podIP
        - name: POD_NAME
          valueFrom:
            fieldRef:
              fieldPath: metadata.name
        - name: POD_NAMESPACE
          valueFrom:
            fieldRef:
              fieldPath: metadata.namespace
        - name: NODE_NAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
        volumeMounts:
        - name: run
          mountPath: /run
        - name: host-cni-conf
          mountPath: /etc/cni/net.d
        - name: flannel-plugin-config
          mountPath: /etc/kube-flannel/
      hostNetwork: true
      nodeSelector:
        beta.kubernetes.io/arch: amd64
      tolerations:
        # Allow the pod to run on the master.  This is required for
        # the master to communicate with pods.
        - key: node-role.kubernetes.io/master
          operator: Exists
          effect: NoSchedule
        # Mark the pod as a critical add-on for rescheduling.
        - key: "CriticalAddonsOnly"
          operator: "Exists"
      volumes:
        - name: run
          hostPath:
            path: /run
        - name: host-cni-conf
          hostPath:
            path: {{ .ConfDir }}
        - name: flannel-plugin-config
          configMap:
            name: flannel-plugin-config-map
        - name: host-cni-bin
          hostPath:
            path: {{ .BinDir }}
  updateStrategy:
    rollingUpdate:
      maxUnavailable: 1
    type: RollingUpdate
`
)
