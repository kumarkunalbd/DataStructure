class Solution {
    public int candy(int[] ratings) {
        int l = 0;
        int r = 0;
        int minT = 1;
        int prevC = 1;
        int upperCap = 0;
        for( r=1; r<ratings.length;r++){
            if(ratings[r]<ratings[r-1]){
                if(prevC == 1){
                    if((r-l)<upperCap){
                        minT += (r-l-1)+1;
                    }else{
                        minT += (r-l)+1;
                    }
                }else{
                    minT +=1;
                }
                prevC = 1;
            }else if(ratings[r]==ratings[r-1]){
                minT +=1;
                l = r;
                prevC = 1;
                upperCap = prevC;
            }else{
                minT += prevC+1;
                prevC = prevC+1;
                l=r;
                upperCap = prevC;
            }
        }
        return minT;
    }
}  