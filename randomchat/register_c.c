#include <stdio.h>
#include <string.h>
#include <sys/socket.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <pthread.h>
#include <sys/types.h>
#include <sys/ipc.h>
#include <errno.h>
#include <stdlib.h>

struct info{
	char id[30];
	char password[30];
};
struct sndd{

	pthread_t *thid1;
	int *sockfd;

};
pthread_t *thid3;
int *sockfd2;
void asciiart();
void _register();
void maaain();
void _login();
void *_wwrite(void *data);
void *_rread(void *data);
void chatconnect();
void maaain2(char *id,char *pw);
void exit_handler1(void);
int main(){
	maaain();
	_register();
	return 0;

}
void maaain2(char *id,char *pw){
	char mid[30];
	char mpw[30];
	int answer;
	strcpy(mid,id);
	strcpy(mpw,pw);
	while(1){

		system("clear");
		asciiart();
		printf("%s님 안녕하세요 ^^\n",mid);
		printf("1.랜덤채팅시작\n");
		printf("2.메인화면이동\n");
		printf("\n번호입력:");
		scanf("%d",&answer);
		__fpurge(stdin);
		if(answer==1){
			chatconnect();
		}else{
			maaain();
		}
		__fpurge(stdin);
	}


}
void exit_handler1(void){

	printf("연결을 끊었습니다. 잠시만 기다리세요\n");
	sleep(10);
	chatconnect();

}
void maaain(){
	int answer;

	while(1){
		system("clear");
asciiart();
printf("1.로그인\n");
printf("2.회원가입\n");
printf("번호입력:");
scanf("%d",&answer);
__fpurge(stdin);
if(answer==2){
_register();
}else if(answer==1){
_login();
}else{
printf("잘못된입력\n");
sleep(2);
}
}	
}
void asciiart(){


	        printf("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\n");
	        printf("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\n");
		printf("░░██████╗  █████╗ ███╗   ██╗██████╗  ██████╗ ███╗   ███╗░░░░░░░░░░░░\n");
		printf("░░██╔══██╗██╔══██╗████╗  ██║██╔══██╗██╔═══██╗████╗ ████║░░░░░░░░░░░░\n");
		printf("░░██████╔╝███████║██╔██╗ ██║██║  ██║██║   ██║██╔████╔██║░░░░░░░░░░░░\n");
		printf("░░██╔══██╗██╔══██║██║╚██╗██║██║  ██║██║   ██║██║╚██╔╝██║░░░░░░░░░░░░\n");
		printf("░░██║  ██║██║  ██║██║ ╚████║██████╔╝╚██████╔╝██║ ╚═╝ ██║░░░░░░░░░░░░\n");
		printf("░░╚═╝  ╚═╝╚═╝  ╚═╝╚═╝  ╚═══╝╚═════╝  ╚═════╝ ╚═╝     ╚═╝░░░░░░░░░░░░\n");
		printf("░░ ██████╗██╗  ██╗ █████╗ ████████╗████████╗██╗███╗   ██╗ ██████╗░░░\n");
		printf("░░██╔════╝██║  ██║██╔══██╗╚══██╔══╝╚══██╔══╝██║████╗  ██║██╔════╝░░░\n");
		printf("░░██║     ███████║███████║   ██║      ██║   ██║██╔██╗ ██║██║  ███╗░░\n");
		printf("░░██║     ██╔══██║██╔══██║   ██║      ██║   ██║██║╚██╗██║██║   ██║░░\n");
		printf("░░╚██████╗██║  ██║██║  ██║   ██║      ██║   ██║██║ ╚████║╚██████╔╝░░\n");
		printf("░░ ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝   ╚═╝      ╚═╝   ╚═╝╚═╝  ╚═══╝ ╚═════╝ ░░\n");
		printf("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\n");
		printf("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░Ver_1░░░░░░░\n");
		printf("░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░░\n\n");



}
void _login(){
int sockfd;
char answer[100];
struct sockaddr_in serveraddr;
struct info r_info;

memset(&serveraddr,0x00,sizeof(serveraddr));
serveraddr.sin_family=AF_INET;
serveraddr.sin_addr.s_addr=inet_addr("127.0.0.1");
serveraddr.sin_port=htons(7891);

sockfd=socket(PF_INET,SOCK_STREAM,IPPROTO_TCP);
connect(sockfd,(struct sockaddr *)&serveraddr,sizeof(serveraddr));

memset(&r_info,0x00,sizeof(r_info));
system("clear");
asciiart();
printf("###로그인###\n");
printf("아이디:");
scanf("%s",&r_info.id);
__fpurge(stdin);
printf("패스워드:");
scanf("%s",&r_info.password);
__fpurge(stdin);
      //strcpy(r_info.id,"admin");
     //strcpy(r_info.password,"1234");
      //strcpy(r_info.nickname,"ad");

   //printf("%s\n",r_info.id);
   //printf("%s\n",r_info.password);

   write(sockfd,(void *)&r_info,sizeof(r_info));
   read(sockfd,answer,100);
   //printf("%s\n",answer);
if(!strcmp(answer,"loginokay")){
	printf("인증되었습니다.\n");
	printf("메뉴 이동중입니다. 잠시만기다려주세요.\n");
	sleep(2);
	maaain2(r_info.id,r_info.password);

}
else{
        printf("아이디 및 비밀번호가 일치하지 않습니다.\n");
        sleep(2);
	close(sockfd);
        maaain();
}
      close(sockfd);

}

void _register(){
	int sockfd;
	char answer[100];
	char answer2[100];
	struct sockaddr_in serveraddr;
	struct info r_info;

	memset(&serveraddr,0x00,sizeof(serveraddr));
	serveraddr.sin_family=AF_INET;
	serveraddr.sin_addr.s_addr=inet_addr("127.0.0.1");
	serveraddr.sin_port=htons(7890);

	sockfd=socket(PF_INET,SOCK_STREAM,IPPROTO_TCP);
	connect(sockfd,(struct sockaddr *)&serveraddr,sizeof(serveraddr));

	memset(&r_info,0x00,sizeof(r_info));
	system("clear");
	asciiart();
	printf("###회원가입###\n");
	printf("아이디를 입력해주세요:");
	scanf("%s",&r_info.id);
__fpurge(stdin);
	printf("패스워드를 입력하세요:");
	scanf("%s",&r_info.password);
__fpurge(stdin);
	//strcpy(r_info.id,"admin");
	//strcpy(r_info.password,"1234");
	//strcpy(r_info.nickname,"ad");
	
	//printf("%s\n",r_info.id);
	//printf("%s\n",r_info.password);

	write(sockfd,(void *)&r_info,sizeof(r_info));
	read(sockfd,answer,100);
	//printf("%s\n",answer);
	if(!strcmp(answer,"okay")){
		printf("등록가능한 아이디입니다.\n");
		printf("정말 등록하시겠습니까(yes/no)?");
		scanf("%s",&answer2);
__fpurge(stdin);
	
		if(!strcmp(answer2,"yes")){
		write(sockfd,"reg",100);
		printf("등록되었습니다.\n");
				sleep(2);
				close(sockfd);
				maaain();
		}
		else if(!strcmp(answer2,"no")){
			printf("취소되었습니다.\n");
				sleep(2);
				close(sockfd);
				maaain();
		}
	        else{
				printf("잘못입력하셨습니다.\n");
				sleep(2);
				close(sockfd);
				maaain();
		}
	}
	else{
		printf("이미존재하는 아이디입니다.\n");
		sleep(2);
		close(sockfd);
		maaain();
	}
	close(sockfd);

}
void chatconnect(){

	int sockfd;
	int signal;
	struct sndd ssdd;
	struct sndd ssdd2;
	struct sockaddr_in serveraddr;
	pthread_t thid1;
	pthread_t thid2;

	atexit(exit_handler1);
	sockfd=socket(PF_INET,SOCK_STREAM,IPPROTO_TCP);
	memset(&serveraddr,0x00,sizeof(serveraddr));
	serveraddr.sin_family=AF_INET;
	serveraddr.sin_addr.s_addr=inet_addr("127.0.0.1");
	serveraddr.sin_port=htons(7892);

	system("clear");
	printf("낯선사람을 찾는중 ..\n");
	connect(sockfd,(struct sockaddr *)&serveraddr,sizeof(serveraddr));

	while(1){
	read(sockfd,&signal,sizeof(signal));
	if(ntohl(signal)==1){
		printf("낯선사람과 연결되었습니다.\n");
		break;
	}
	}
	

	
	pthread_create(&thid1,NULL,_wwrite,(void *)&sockfd);
	ssdd.sockfd=&sockfd;
	ssdd.thid1=&thid1;
	pthread_create(&thid2,NULL,_rread,(void *)&ssdd);
	thid3=&thid2;
	pthread_join(thid2,NULL);
	pthread_join(thid1,NULL);
	printf("낯선사람과의 연결이 끊겼습니다.(잠시만요)\n");
	sleep(3);
	chatconnect();

}

void *_wwrite(void *data){
int *sockfd=(int *)data;
char buffer[1024];
while(1){
gets(buffer);
__fpurge(stdin);
if(!strcmp(buffer,"reconnect")){
		close(*sockfd);
		pthread_cancel(*thid3);
		exit(0);
		}

send(*sockfd,buffer,1024,0);
}
}

void *_rread(void *data){
struct sndd *ssdd;
ssdd=(struct sndd *)data;
char buffer[1024];
int *sockfd=ssdd->sockfd;
pthread_t *pth=ssdd->thid1;
//ssdd=(struct sndd *)&data;
while(1){
memset(buffer,0x00,1024);
if(recv(*sockfd,buffer,1024,0)<=0){
	fprintf(stderr,"Recv error:%s\n",strerror(errno));
	break;
}
if(!strcmp(buffer,"close")){
	close(*sockfd);
	pthread_cancel(*pth);
	break;
}
printf("낯선사람:%s\n",buffer);
}
}
