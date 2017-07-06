package gosrt
// #cgo pkg-config: srt
// #include <srt/srt.h>
import "C"
const (
    OPT_MSS =             C.SRTO_MSS           
    OPT_SNDSYN =          C.SRTO_SNDSYN        
    OPT_RCVSYN =          C.SRTO_RCVSYN        
    OPT_CC =              C.SRTO_CC            
    OPT_FC =              C.SRTO_FC            
    OPT_SNDBUF =          C.SRTO_SNDBUF        
    OPT_RCVBUF =          C.SRTO_RCVBUF        
    OPT_LINGER =          C.SRTO_LINGER        
    OPT_UDP_SNDBUF =      C.SRTO_UDP_SNDBUF    
    OPT_UDP_RCVBUF =      C.SRTO_UDP_RCVBUF    
    OPT_MAXMSG =          C.SRTO_MAXMSG        
    OPT_MSGTTL =          C.SRTO_MSGTTL        
    OPT_RENDEZVOUS =      C.SRTO_RENDEZVOUS    
    OPT_SNDTIMEO =        C.SRTO_SNDTIMEO      
    OPT_RCVTIMEO =        C.SRTO_RCVTIMEO      
    OPT_REUSEADDR =       C.SRTO_REUSEADDR     
    OPT_MAXBW =           C.SRTO_MAXBW         
    OPT_STATE =           C.SRTO_STATE         
    OPT_EVENT =           C.SRTO_EVENT         
    OPT_SNDDATA =         C.SRTO_SNDDATA       
    OPT_RCVDATA =         C.SRTO_RCVDATA       
    OPT_SENDER =          C.SRTO_SENDER   
    OPT_TSBPDMODE =       C.SRTO_TSBPDMODE
    OPT_TSBPDDELAY =      C.SRTO_TSBPDDELAY
    OPT_LATENCY =         C.SRTO_LATENCY  
    OPT_INPUTBW =         C.SRTO_INPUTBW  
    OPT_OHEADBW =         C.SRTO_OHEADBW       
    OPT_PASSPHRASE =      C.SRTO_PASSPHRASE
    OPT_PBKEYLEN =        C.SRTO_PBKEYLEN      
    OPT_KMSTATE =         C.SRTO_KMSTATE       
    OPT_IPTTL =           C.SRTO_IPTTL    
    OPT_IPTOS =           C.SRTO_IPTOS         
    OPT_TLPKTDROP =       C.SRTO_TLPKTDROP
    OPT_TSBPDMAXLAG =     C.SRTO_TSBPDMAXLAG   
    OPT_NAKREPORT =       C.SRTO_NAKREPORT
    OPT_VERSION =         C.SRTO_VERSION  
    OPT_PEERVERSION =     C.SRTO_PEERVERSION   
    OPT_CONNTIMEO =       C.SRTO_CONNTIMEO
    OPT_TWOWAYDATA =      C.SRTO_TWOWAYDATA
    OPT_SNDPBKEYLEN =     C.SRTO_SNDPBKEYLEN 
    OPT_RCVPBKEYLEN =     C.SRTO_RCVPBKEYLEN
    OPT_SNDPEERKMSTATE =  C.SRTO_SNDPEERKMSTATE
    OPT_RCVKMSTATE =      C.SRTO_RCVKMSTATE
    OPT_LOSSMAXTTL =      C.SRTO_LOSSMAXTTL
)
