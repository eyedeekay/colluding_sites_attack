package echosam

import (
	"fmt"
	//"path/filepath"
	"strconv"
)

//Option is a EchoSAM Option
type Option func(*EchoSAM) error

//SetFilePath sets the path to save the config file at.
func SetFilePath(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.SAMForwarder.Config().FilePath = s
		return nil
	}
}

//SetType sets the type of the forwarder server
func SetType(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if s == "http" {
			c.SAMForwarder.Config().Type = s
			return nil
		} else {
			c.SAMForwarder.Config().Type = "http"
			return nil
		}
	}
}

//SetSigType sets the type of the forwarder server
func SetSigType(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if s == "" {
			c.SAMForwarder.Config().SigType = ""
		} else if s == "DSA_SHA1" {
			c.SAMForwarder.Config().SigType = "DSA_SHA1"
		} else if s == "ECDSA_SHA256_P256" {
			c.SAMForwarder.Config().SigType = "ECDSA_SHA256_P256"
		} else if s == "ECDSA_SHA384_P384" {
			c.SAMForwarder.Config().SigType = "ECDSA_SHA384_P384"
		} else if s == "ECDSA_SHA512_P521" {
			c.SAMForwarder.Config().SigType = "ECDSA_SHA512_P521"
		} else if s == "EdDSA_SHA512_Ed25519" {
			c.SAMForwarder.Config().SigType = "EdDSA_SHA512_Ed25519"
		} else {
			c.SAMForwarder.Config().SigType = "EdDSA_SHA512_Ed25519"
		}
		return nil
	}
}

//SetSaveFile tells the router to save the tunnel's keys long-term
func SetSaveFile(b bool) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.SAMForwarder.Config().SaveFile = b
		return nil
	}
}

//SetHost sets the host of the service to forward
func SetHost(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.SAMForwarder.Config().TargetHost = s
		return nil
	}
}

//SetPort sets the port of the service to forward
func SetPort(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid TCP Server Target Port %s; non-number ", s)
		}
		if port < 65536 && port > -1 {
			c.SAMForwarder.Config().TargetPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetSAMHost sets the host of the EchoSAM's SAM bridge
func SetSAMHost(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.SAMForwarder.Config().SamHost = s
		return nil
	}
}

//SetSAMPort sets the port of the EchoSAM's SAM bridge using a string
func SetSAMPort(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		port, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("Invalid SAM Port %s; non-number", s)
		}
		if port < 65536 && port > -1 {
			c.SAMForwarder.Config().SamPort = s
			return nil
		}
		return fmt.Errorf("Invalid port")
	}
}

//SetName sets the host of the EchoSAM's SAM bridge
func SetName(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.SAMForwarder.Config().TunName = s
		return nil
	}
}

//SetInLength sets the number of hops inbound
func SetInLength(u int) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if u < 7 && u >= 0 {
			c.SAMForwarder.Config().InLength = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutLength sets the number of hops outbound
func SetOutLength(u int) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if u < 7 && u >= 0 {
			c.SAMForwarder.Config().OutLength = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel length")
	}
}

//SetInVariance sets the variance of a number of hops inbound
func SetInVariance(i int) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if i < 7 && i > -7 {
			c.SAMForwarder.Config().InVariance = i
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel length")
	}
}

//SetOutVariance sets the variance of a number of hops outbound
func SetOutVariance(i int) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if i < 7 && i > -7 {
			c.SAMForwarder.Config().OutVariance = i
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel variance")
	}
}

//SetInQuantity sets the inbound tunnel quantity
func SetInQuantity(u int) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if u <= 16 && u > 0 {
			c.SAMForwarder.Config().InQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel quantity")
	}
}

//SetOutQuantity sets the outbound tunnel quantity
func SetOutQuantity(u int) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if u <= 16 && u > 0 {
			c.SAMForwarder.Config().OutQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel quantity")
	}
}

//SetInBackups sets the inbound tunnel backups
func SetInBackups(u int) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if u < 6 && u >= 0 {
			c.SAMForwarder.Config().InBackupQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid inbound tunnel backup quantity")
	}
}

//SetOutBackups sets the inbound tunnel backups
func SetOutBackups(u int) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if u < 6 && u >= 0 {
			c.SAMForwarder.Config().OutBackupQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid outbound tunnel backup quantity")
	}
}

//SetEncrypt tells the router to use an encrypted leaseset
func SetEncrypt(b bool) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if b {
			c.SAMForwarder.Config().EncryptLeaseSet = true
			return nil
		}
		c.SAMForwarder.Config().EncryptLeaseSet = false
		return nil
	}
}

//SetServeDir sets the path to the directory you want to serve
/*func SetServeDir(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		var err error
		c.ServeDir, err = filepath.Abs(s)
		if err != nil {
			return err
		}
		return nil
	}
}*/

//SetLeaseSetKey sets the host of the EchoSAM's SAM bridge
func SetLeaseSetKey(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.SAMForwarder.Config().LeaseSetKey = s
		return nil
	}
}

//SetLeaseSetPrivateKey sets the host of the EchoSAM's SAM bridge
func SetLeaseSetPrivateKey(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.SAMForwarder.Config().LeaseSetPrivateKey = s
		return nil
	}
}

//SetLeaseSetPrivateSigningKey sets the host of the EchoSAM's SAM bridge
func SetLeaseSetPrivateSigningKey(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.SAMForwarder.Config().LeaseSetPrivateSigningKey = s
		return nil
	}
}

//SetMessageReliability sets the host of the EchoSAM's SAM bridge
func SetMessageReliability(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.SAMForwarder.Config().MessageReliability = s
		return nil
	}
}

//SetAllowZeroIn tells the tunnel to accept zero-hop peers
func SetAllowZeroIn(b bool) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if b {
			c.SAMForwarder.Config().InAllowZeroHop = true
			return nil
		}
		c.SAMForwarder.Config().InAllowZeroHop = false
		return nil
	}
}

//SetAllowZeroOut tells the tunnel to accept zero-hop peers
func SetAllowZeroOut(b bool) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if b {
			c.SAMForwarder.Config().OutAllowZeroHop = true
			return nil
		}
		c.SAMForwarder.Config().OutAllowZeroHop = false
		return nil
	}
}

//SetCompress tells clients to use compression
func SetCompress(b bool) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if b {
			c.SAMForwarder.Config().UseCompression = true
			return nil
		}
		c.SAMForwarder.Config().UseCompression = false
		return nil
	}
}

//SetFastRecieve tells clients to use compression
func SetFastRecieve(b bool) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if b {
			c.SAMForwarder.Config().FastRecieve = true
			return nil
		}
		c.SAMForwarder.Config().FastRecieve = false
		return nil
	}
}

//SetReduceIdle tells the connection to reduce it's tunnels during extended idle time.
func SetReduceIdle(b bool) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if b {
			c.SAMForwarder.Config().ReduceIdle = true
			return nil
		}
		c.SAMForwarder.Config().ReduceIdle = false
		return nil
	}
}

//SetReduceIdleTime sets the time to wait before reducing tunnels to idle levels
func SetReduceIdleTime(u int) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.SAMForwarder.Config().ReduceIdleTime = 300000
		if u >= 6 {
			c.SAMForwarder.Config().ReduceIdleTime = (u * 60) * 1000
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in minutes) %v", u)
	}
}

//SetReduceIdleTimeMs sets the time to wait before reducing tunnels to idle levels in milliseconds
func SetReduceIdleTimeMs(u int) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.SAMForwarder.Config().ReduceIdleTime = 300000
		if u >= 300000 {
			c.SAMForwarder.Config().ReduceIdleTime = u
			return nil
		}
		return fmt.Errorf("Invalid reduce idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetReduceIdleQuantity sets minimum number of tunnels to reduce to during idle time
func SetReduceIdleQuantity(u int) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if u < 5 {
			c.SAMForwarder.Config().ReduceIdleQuantity = u
			return nil
		}
		return fmt.Errorf("Invalid reduce tunnel quantity")
	}
}

//SetCloseIdle tells the connection to close it's tunnels during extended idle time.
func SetCloseIdle(b bool) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if b {
			c.SAMForwarder.Config().CloseIdle = true
			return nil
		}
		c.SAMForwarder.Config().CloseIdle = false
		return nil
	}
}

//SetCloseIdleTime sets the time to wait before closing tunnels to idle levels
func SetCloseIdleTime(u int) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.SAMForwarder.Config().CloseIdleTime = 300000
		if u >= 6 {
			c.SAMForwarder.Config().CloseIdleTime = (u * 60) * 1000
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in minutes) %v", u)
	}
}

//SetCloseIdleTimeMs sets the time to wait before closing tunnels to idle levels in milliseconds
func SetCloseIdleTimeMs(u int) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.SAMForwarder.Config().CloseIdleTime = 300000
		if u >= 300000 {
			c.SAMForwarder.Config().CloseIdleTime = u
			return nil
		}
		return fmt.Errorf("Invalid close idle timeout(Measured in milliseconds) %v", u)
	}
}

//SetAccessListType tells the system to treat the accessList as a whitelist
func SetAccessListType(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if s == "whitelist" {
			c.SAMForwarder.Config().AccessListType = "whitelist"
			return nil
		} else if s == "blacklist" {
			c.SAMForwarder.Config().AccessListType = "blacklist"
			return nil
		} else if s == "none" {
			c.SAMForwarder.Config().AccessListType = ""
			return nil
		} else if s == "" {
			c.SAMForwarder.Config().AccessListType = ""
			return nil
		}
		return fmt.Errorf("Invalid Access list type(whitelist, blacklist, none)")
	}
}

//SetAccessList tells the system to treat the accessList as a whitelist
func SetAccessList(s []string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		if len(s) > 0 {
			for _, a := range s {
				c.SAMForwarder.Config().AccessList = append(c.SAMForwarder.Config().AccessList, a)
			}
			return nil
		}
		return nil
	}
}

//SetKeyFile sets
func SetKeyFile(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.SAMForwarder.Config().KeyFilePath = s
		return nil
	}
}

//SetCSS sets the path to save the config file at.
func SetCSS(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.CSS = s
		return nil
	}
}

//SetFilePath sets the path to save the config file at.
func SetFingerJS(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.FingerprintJS = s
		return nil
	}
}

//SetFilePath sets the path to save the config file at.
func SetLocalJS(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.LocalJS = s
		return nil
	}
}

//SetFingerFile sets the path to save the config file at.
func SetFingerFile(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.FingerFile = s
		return nil
	}
}

//SetWebRTC sets the path to save the config file at.
func SetWebRTC(s string) func(*EchoSAM) error {
	return func(c *EchoSAM) error {
		c.RTCFile = s
		return nil
	}
}
