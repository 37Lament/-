#include "bits/stdc++.h"
using namespace std;
struct {
    bool sl;
    long room;
}g[10005][10000];
int search (int x,int y,int m){
    long temp=g[x][y].room;
    while(1){
        if(g[x][y].sl==1) temp--;
        if(temp==0) break;
        if(y<m-1) y++;
        else y=0;}
        return y;
}

int main(){
    int n,m,start;
    long pw=0;
    cin>>n>>m;
    for (int i = 1; i <=n ; i++)
    {for(int j=0;j<m;j++){
            cin>>g[i][j].sl>>g[i][j].room;
        }}
    cin>>start;
    for(int i=1;i<=n;i++){
        pw+=g[i][start].room;
        start=search(i,start,m);
    }
    cout<<pw<<endl;
    return 0;
}
