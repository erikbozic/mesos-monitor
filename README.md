# mesos-monitor
**Work in progress!**    
Command line tool for monitoring mesos task logs

# Usage
Run binary with flag `-m` or `-master` and pass in the http url for mesos master.  
`` ./mesos-monitor -m http://localhost:5050 ``  
This will get all tasks known to master and let you specify tasks which you want to monitor logs from.
