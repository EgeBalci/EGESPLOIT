# EGESPLOIT [![License](https://img.shields.io/github/license/mashape/apistatus.svg?maxAge=2592000)](https://raw.githubusercontent.com/EgeBalci/EGESPLOIT/master/LICENSE?token=AQYjCTVzZPDeUfD7VWrNhzERQAM-rMYZks5Xle7twA%3D%3D)  [![Donate](https://img.shields.io/badge/Donate-Patreon-green.svg)](http://patreon.com/user?u=3556027)


EGESPLOIT is a golang library for malware development, it has few unique functions for meterpreter integration.


#DOCUMENTATION


            CalculateChecksum(x) : Function calculates x digit 8 bit checksum for reverse HTTP/HTTPS meterpreter connections, returns the calculated checksum as string.
            
            Meterpreter(ConType, Address) : Function launches a meterpreter connection, takes 2 parameters connection type (HTTP/HTTPS/TCP) and Address (127.0.0.1:4444), function returns a string for error handling.
            
            Persistence() : Function copys and adds the running binary to startup registry.
            
            Sysguide() : Function returns the current directory, running OS version, username, antivirus name as strings.
            
            Keylogger(LOGS) : Function takes a string pointer as parameter and starts a keylogger,all key logs are saved at given parameter.
            
            Please(Command) : Function executes the given parameter with runas command. (Asks permission for higher level operations)  
            
            BypassAV() : Function bypasses the anti virus heroustic detections, takes a integer as parameter for defining the intensity level.
            
            Dispatch(Base64_Binary,BinaryName, Parameters) : Function drops a binary and executes it, takes tree strings as parameter base64 encoded binary, binary name and parameters.
            
            Distract() : Functions execute a forkbomb bat file for distracting the user.
            
            Dos() : Function start a dos atack to given target (http://example.com)
            
            SyscallExecute(Shellcode) : Function executes the given shellcode(byte array) with system call.
            
            ThreadExecute(Shellcode) : Function executes the given shellcode(byte array) with CreateThread function.

            WifiList() : Functions returns he wifi connection history.
            
            #RSE#
            RSE stands for "Reduced Sized Exploits", functions under RSE folder are build with windows api calls for reducing payload sizes.
      
      
![](http://i.imgur.com/8L1wmjo.png)

   ![](http://i.imgur.com/N2bhpR9.jpg)

Bitcoin: 16GvMV7eZH22p4rLQuu8h2gbgSLYr11KBM
