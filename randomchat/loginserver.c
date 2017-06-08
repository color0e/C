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



int idpwconfirm(char *id,char *pw);

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
	serveraddr.sin_port=htons(7891);


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
	confirm=idpwconfirm(r_info.id,r_info.password);
	if(confirm==-1){
		write(clientsock,"loginnono",100);
	}
	else{
		write(clientsock,"loginokay",100);
	}
	close(clientsock);
	}
	else if(pid>0){
	close(clientsock);
	}

	}
}

int idpwconfirm(char *id,char *pw){
int fd=open("/var/cpasswd",O_RDONLY);
char cid[10];
char cpw[10];
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
lseek(fd,-11,SEEK_CUR);
read(fd,cpw,10);
if(strcmp(cpw,pw)!=0){
	close(fd);
return -1;
}
close(fd);
return 0;
}
