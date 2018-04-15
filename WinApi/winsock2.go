package WinApi

import (
	"syscall"
	"unsafe"
)

var (
	winSock2dll                         syscall.Handle
	fnioctlsocket                       uintptr
)



const(
	IPPROTO_IP   = 0; // dummy for IP
	IPPROTO_HOPOPTS = 0; // IPv6 hop-by-hop options
	IPPROTO_ICMP = 1; // control message protocol
	IPPROTO_IGMP = 2; // internet group management protocol
	IPPROTO_GGP  = 3; // gateway^2 (deprecated)
	IPPROTO_IPV4 = 4; // IPv4
	IPPROTO_ST   = 5;
	IPPROTO_TCP  = 6; // tcp
	IPPROTO_CBT  = 7;
	IPPROTO_EGP  = 8;
	IPPROTO_IGP  = 9;
	IPPROTO_PUP  = 12; // pup
	IPPROTO_UDP  = 17; // user datagram protocol
	IPPROTO_IDP  = 22; // xns idp
	IPPROTO_RDP  = 27;
	IPPROTO_IPV6 = 41; // IPv6
	IPPROTO_ROUTING        = 43;              // IPv6 routing header
	IPPROTO_FRAGMENT       = 44;              // IPv6 fragmentation header
	IPPROTO_ESP            = 50;              // IPsec ESP header
	IPPROTO_AH             = 51;              // IPsec AH
	IPPROTO_ICMPV6         = 58;              // ICMPv6
	IPPROTO_NONE           = 59;              // IPv6 no next header
	IPPROTO_DSTOPTS        = 60;              // IPv6 destination options
	IPPROTO_ND   = 77; // UNOFFICIAL net disk proto
	IPPROTO_ICLFXBM = 78;
	IPPROTO_PIM     = 103;
	IPPROTO_PGM     = 113;
	IPPROTO_L2TP    = 115;
	IPPROTO_SCTP    = 132;

	IPPROTO_RAW  = 255; // raw IP packet
	IPPROTO_MAX  = 256;
	//
	//  These are reserved for internal use by Windows.
	//
	IPPROTO_RESERVED_RAW  = 257;
	IPPROTO_RESERVED_IPSEC  = 258;
	IPPROTO_RESERVED_IPSECOFFLOAD  = 259;
	IPPROTO_RESERVED_MAX  = 260;

	//
	// Port/socket numbers: network standard functions
	//

	IPPORT_TCPMUX     = 1;
	IPPORT_ECHO       = 7;
	IPPORT_DISCARD    = 9;
	IPPORT_SYSTAT     = 11;
	IPPORT_DAYTIME    = 13;
	IPPORT_NETSTAT    = 15;
	IPPORT_QOTD       = 17;
	IPPORT_MSP        = 18;
	IPPORT_CHARGEN    = 19;
	IPPORT_FTP_DATA   = 20;
	IPPORT_FTP        = 21;
	IPPORT_TELNET     = 23;
	IPPORT_SMTP       = 25;
	IPPORT_TIMESERVER = 37;
	IPPORT_NAMESERVER = 42;
	IPPORT_WHOIS      = 43;
	IPPORT_MTP        = 57;

	//
	// Port/socket numbers: host specific functions
	//

	IPPORT_TFTP    = 69;
	IPPORT_RJE     = 77;
	IPPORT_FINGER  = 79;
	IPPORT_TTYLINK = 87;
	IPPORT_SUPDUP  = 95;

	//
	// UNIX TCP sockets
	//
	IPPORT_POP3         = 110;
	IPPORT_NTP          = 123;
	IPPORT_EPMAP        = 135;
	IPPORT_NETBIOS_NS   = 137;
	IPPORT_NETBIOS_DGM  = 138;
	IPPORT_NETBIOS_SSN  = 139;
	IPPORT_IMAP         = 143;
	IPPORT_SNMP         = 161;
	IPPORT_SNMP_TRAP    = 162;
	IPPORT_IMAP3        = 220;
	IPPORT_LDAP         = 389;
	IPPORT_HTTPS        = 443;
	IPPORT_MICROSOFT_DS = 445;

	IPPORT_EXECSERVER  = 512;
	IPPORT_LOGINSERVER = 513;
	IPPORT_CMDSERVER   = 514;
	IPPORT_EFSSERVER   = 520;

	//
	// UNIX UDP sockets
	//

	IPPORT_BIFFUDP     = 512;
	IPPORT_WHOSERVER   = 513;
	IPPORT_ROUTESERVER = 520;

	// 520+1 also used

	//
	// Ports < IPPORT_RESERVED are reserved for
	// privileged processes (e.g. root).
	//

	IPPORT_RESERVED = 1024;

	IPPORT_REGISTERED_MIN = IPPORT_RESERVED;
	IPPORT_REGISTERED_MAX = 0xbfff;
	IPPORT_DYNAMIC_MIN    = 0xc000;
	IPPORT_DYNAMIC_MAX    = 0xffff;


	FIONBIO = syscall.IOC_IN | ((4 & 0x7f) << 16) | (102 << 8) | 126 // set/clear non-blocking i/o

	IP_OPTIONS = 1
	IP_HDRINCL = 2
	IP_TOS = 3
	IP_TTL = 4
	IP_MULTICAST_IF = 9

)

func init()  {
	winSock2dll, _ = syscall.LoadLibrary("ws2_32.dll")
	fnioctlsocket,_ = syscall.GetProcAddress(winSock2dll, "ioctlsocket")
}


func Ioctlsocket(sockHandle syscall.Handle,cmd uint32,arpg *uint32)int  {
	r1, _, _ := syscall.Syscall(fnioctlsocket,3,uintptr(sockHandle),uintptr(cmd),uintptr(unsafe.Pointer(arpg)))
	return int(r1)
}
