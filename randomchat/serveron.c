#include <stdio.h>
#include <unistd.h>

int main(){

	pid_t pid;

	pid=fork();
	if(pid==0){
		pid=fork();
		if(pid==0){
		execlp("./r_server","./r_server",NULL);
		
		}else{
		execlp("./chatserver","./chatserver",NULL);

		}
	}else{
		execlp("./loginserver","./loginserver",NULL);
	}

	return 0;
}
