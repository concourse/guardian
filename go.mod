module github.com/concourse/guardian

require (
	code.cloudfoundry.org/archiver v0.0.0-20180525162158-e135af3d5a2a
	code.cloudfoundry.org/clock v0.0.0-20180518195852-02e53af36e6c
	code.cloudfoundry.org/commandrunner v0.0.0-20180212143422-501fd662150b
	code.cloudfoundry.org/debugserver v0.0.0-20170501225606-70715da12ee9
	code.cloudfoundry.org/garden v0.0.0-20180820151144-7999b305fe99
	code.cloudfoundry.org/grootfs v0.0.0-20180525140952-c01568707fea
	code.cloudfoundry.org/idmapper v0.0.0-20170928154045-bd44efed5494
	code.cloudfoundry.org/lager v1.1.0
	code.cloudfoundry.org/localip v0.0.0-20170223024724-b88ad0dea95c
	github.com/BurntSushi/toml v0.3.0 // indirect
	github.com/Microsoft/go-winio v0.4.7 // indirect
	github.com/Microsoft/hcsshim v0.6.11 // indirect
	github.com/apoydence/eachers v0.0.0-20161009051034-dc08b96c9060 // indirect
	github.com/beorn7/perks v0.0.0-20180321164747-3a771d992973 // indirect
	github.com/bmizerany/pat v0.0.0-20170815010413-6226ea591a40 // indirect
	github.com/boltdb/bolt v1.3.1 // indirect
	github.com/burntsushi/toml v0.3.0
	github.com/cloudfoundry/dropsonde v1.0.0
	github.com/cloudfoundry/gosigar v1.1.0
	github.com/cloudfoundry/gosteno v0.0.0-20150423193413-0c8581caea35 // indirect
	github.com/cloudfoundry/loggregatorlib v0.0.0-20170823162133-36eddf15ef12 // indirect
	github.com/cloudfoundry/sonde-go v0.0.0-20171206171820-b33733203bb4 // indirect
	github.com/containerd/cgroups v0.0.0-20180820144417-3024bc7cc0c8 // indirect
	github.com/containerd/console v0.0.0-20180307192801-cb7008ab3d83 // indirect
	github.com/containerd/containerd v1.1.0
	github.com/containerd/continuity v0.0.0-20180522172239-2d3749b4da56 // indirect
	github.com/containerd/fifo v0.0.0-20180307165137-3d5202aec260 // indirect
	github.com/containerd/go-runc v0.0.0-20180511165015-301f7c1fbbc3 // indirect
	github.com/containerd/typeurl v0.0.0-20170912152501-f6943554a7e7 // indirect
	github.com/coreos/go-systemd v0.0.0-20180705093442-88bfeed483d3 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/docker v0.0.0-20180412184420-05346355dbd4
	github.com/docker/go-connections v0.3.0 // indirect
	github.com/docker/go-events v0.0.0-20170721190031-9461782956ad // indirect
	github.com/docker/go-metrics v0.0.0-20180209012529-399ea8c73916 // indirect
	github.com/docker/go-units v0.3.3 // indirect
	github.com/eapache/go-resiliency v1.1.0
	github.com/fsnotify/fsnotify v1.4.7 // indirect
	github.com/godbus/dbus v4.1.0+incompatible // indirect
	github.com/gogo/googleapis v1.0.0 // indirect
	github.com/gogo/protobuf v1.0.0 // indirect
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b // indirect
	github.com/golang/protobuf v1.1.0 // indirect
	github.com/google/go-cmp v0.2.0 // indirect
	github.com/gotestyourself/gotestyourself v2.1.0+incompatible // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/hashicorp/errwrap v0.0.0-20141028054710-7554cd9344ce // indirect
	github.com/hashicorp/go-multierror v0.0.0-20171204182908-b7773ae21874
	github.com/hpcloud/tail v1.0.0 // indirect
	github.com/jessevdk/go-flags v1.4.0
	github.com/kardianos/osext v0.0.0-20170510131534-ae77be60afb1
	github.com/mailru/easyjson v0.0.0-20180323154445-8b799c424f57 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/copystructure v0.0.0-20170525013902-d23ffcb85de3
	github.com/mitchellh/go-ps v0.0.0-20170309133038-4fdf99ab2936
	github.com/mitchellh/reflectwalk v0.0.0-20170726202117-63d60e9d0dbc // indirect
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d
	github.com/onsi/ginkgo v1.6.0
	github.com/onsi/gomega v1.4.1
	github.com/opencontainers/go-digest v1.0.0-rc1
	github.com/opencontainers/image-spec v1.0.1
	github.com/opencontainers/runc v0.0.0-20180322154912-ec9bf5058614
	github.com/opencontainers/runtime-spec v1.0.1
	github.com/opencontainers/selinux v1.0.0-rc1 // indirect
	github.com/pkg/errors v0.8.0
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v0.8.0 // indirect
	github.com/prometheus/client_model v0.0.0-20180712105110-5c3871d89910 // indirect
	github.com/prometheus/common v0.0.0-20180801064454-c7de2306084e // indirect
	github.com/prometheus/procfs v0.0.0-20180725123919-05ee40e3a273 // indirect
	github.com/sirupsen/logrus v1.0.5
	github.com/st3v/glager v0.3.0
	github.com/stretchr/testify v1.2.2 // indirect
	github.com/syndtr/gocapability v0.0.0-20180223013746-33e07d32887e // indirect
	github.com/tedsuo/ifrit v0.0.0-20180410193936-e89a512c3162
	github.com/tedsuo/rata v1.0.0 // indirect
	github.com/tscolari/lagregator v0.0.0-20161103133944-b0fb43b01861
	github.com/urfave/cli v1.20.0
	github.com/vbatts/tar-split v0.10.2 // indirect
	github.com/vishvananda/netlink v1.0.0
	github.com/vishvananda/netns v0.0.0-20171111001504-be1fbeda1936
	golang.org/x/crypto v0.0.0-20180515001509-1a580b3eff78 // indirect
	golang.org/x/net v0.0.0-20180522190444-9ef9f5bb98a1 // indirect
	golang.org/x/sync v0.0.0-20180314180146-1d60e4601c6f // indirect
	golang.org/x/sys v0.0.0-20180525142821-c11f84a56e43
	golang.org/x/text v0.3.0 // indirect
	google.golang.org/genproto v0.0.0-20180518175338-11a468237815 // indirect
	google.golang.org/grpc v1.12.0 // indirect
	gopkg.in/airbrake/gobrake.v2 v2.0.9 // indirect
	gopkg.in/fsnotify.v1 v1.4.7 // indirect
	gopkg.in/gemnasium/logrus-airbrake-hook.v2 v2.1.2 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.2.1
	gotest.tools v2.1.0+incompatible // indirect
)
