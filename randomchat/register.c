#include <stdio.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <string.h>
#include <fcntl.h>
struct info{
	char id[30];
	char password[30];
};



int idconfirm(char *id);
int _register(char id[10],char password[10]);

int main(){

	pid_t pid;
	char *id;
	char *password;
	char *nickname;
	char answer2[100];
	int serversock,clientsock,clientlen,confirm;
	struct sockaddr_in serveraddr,clientaddr;
	struct info r_info;
	
	memset(&serveraddr,0x00,sizeof(serveraddr));
	memset(&serveraddr,0x00,sizeof(clientaddr));
	serveraddr.sin_family=AF_INET;
	serveraddr.sin_addr.s_addr=htonl(INADDR_ANY);
	serveraddr.sin_port=htons(7890);


	serversock=socket(PF_INET,SOCK_STREAM,IPPROTO_TCP);
	if(serversock==-1){
		printf("serversock create failed .. \n");
		return 1;
	}
	bind(serversock,(struct sockaddr *)&serveraddr,sizeof(serveraddr));
	listen(serversock,5);
	clientlen=sizeof(clientaddr);

	memset(&r_info,0x00,sizeof(r_info));
	while(1){

	clientsock=accept(serversock,(struct sockaddr *)&clientaddr,&clientlen);
	if(clientsock==-1){
		printf("accept error\n");
		break;
	}
	pid=fork();
	if(pid==0){
	read(clientsock,(void *)&r_info,sizeof(r_info));
	printf("%s\n",r_info.id);
	printf("%s\n",r_info.password);

	//id confirm
	confirm=idconfirm(r_info.id);
	if(confirm==-1){
		write(clientsock,"okay",100);
		read(clientsock,&answer2,100);
		if(!strcmp(answer2,"reg")){
			if(_register(r_info.id,r_info.password)==-1){
			printf("error\n");
			}else{
			printf("등록완료\n");
			}
		}
		else{
			printf("등록취소\n");
		}
	}
	else{
		write(clientsock,"nono",100);
	}
	close(clientsock);
	}
	else if(pid>0){
	close(clientsock);
	}

	}
}

int idconfirm(char *id){
int fd=open("/var/cpasswd",O_RDONLY);
char cid[10];
unsigned char byte;
int length;
if(fd==-1){
	printf("read error\n");
	return 1;
}
while((unsigned int)(strcmp(cid,id))){
	char pass[10];
	if(read(fd,cid,10)!=10){
		close(fd);
		return -1;
	}
	if(read(fd,&byte,1)!=1){
		close(fd);
		return -1;
	}
	if(read(fd,pass,10)!=10){
		close(fd);
		return -1;
	}
	if(read(fd,&byte,1)!=1){
		close(fd);
		return -1;
	}
}
close(fd);
return 0;
}
int _register(char id[10],char pass[10]){
int fd;
fd=open("/var/cpasswd",O_WRONLY|O_CREAT|O_APPEND,S_IRUSR|S_IWUSR);
if(fd==-1)
	return -1;

if(write(fd,id,10)==-1){
	return -1;
}
write(fd,"\n",1);
write(fd,pass,10);
write(fd,"\n",1);
if(close(fd)==-1)
	return -1;

	return 0;

}













