package echosam

import (
	"log"
	"net/http"

	"github.com/eyedeekay/sam-forwarder/interface"
	"github.com/eyedeekay/sam-forwarder/tcp"
)

//EchoSAM is a structure which automatically configured the forwarding of
//a local service to i2p over the SAM API.
type EchoSAM struct {
	*samforwarder.SAMForwarder
	FingerprintJS string
	FingerFile    string
	LocalJS       string
	CSS           string
	up            bool
}

var err error

func (f *EchoSAM) GetType() string {
	return "echosam"
}

func (f *EchoSAM) ServeParent() {
	log.Println("Starting eepsite server", f.Base32())
	if err = f.SAMForwarder.Serve(); err != nil {
		f.Cleanup()
	}
}

//Serve starts the SAM connection and and forwards the local host:port to i2p
func (f *EchoSAM) Serve() error {
	go f.ServeParent()
	if f.Up() {
		log.Println("Starting web server", f.Target())
		if err := http.ListenAndServe(f.Target(), f); err != nil {
			return err
		}
	}
	return nil
}

func (f *EchoSAM) Up() bool {
	return f.up
}

//Close shuts the whole thing down.
func (f *EchoSAM) Close() error {
	return f.SAMForwarder.Close()
}

func (s *EchoSAM) Load() (samtunnel.SAMTunnel, error) {
	if !s.up {
		log.Println("Started putting tunnel up")
	}
	f, e := s.SAMForwarder.Load()
	if e != nil {
		return nil, e
	}
	s.SAMForwarder = f.(*samforwarder.SAMForwarder)
	s.up = true
	log.Println("Finished putting tunnel up")
	return s, nil
}

//NewEchoSAM makes a new SAM forwarder with default options, accepts host:port arguments
func NewEchoSAM(host, port string) (*EchoSAM, error) {
	return NewEchoSAMFromOptions(SetHost(host), SetPort(port))
}

//NewEchoSAMFromOptions makes a new SAM forwarder with default options, accepts host:port arguments
func NewEchoSAMFromOptions(opts ...func(*EchoSAM) error) (*EchoSAM, error) {
	var s EchoSAM
	s.SAMForwarder = &samforwarder.SAMForwarder{}
	log.Println("Initializing eephttpd")
	for _, o := range opts {
		if err := o(&s); err != nil {
			return nil, err
		}
	}
	s.SAMForwarder.Config().SaveFile = true
	l, e := s.Load()
	//log.Println("Options loaded", s.Print())
	if e != nil {
		return nil, e
	}
	return l.(*EchoSAM), nil
}
