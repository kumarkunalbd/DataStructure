class Solution {
    public int reversePairs(int[] nums) {
        long[] numsL = new long[nums.length];
        for(int i=0; i<nums.length;i++){
            numsL[i] = nums[i];
        }
        //System.out.printf("Intial=%s\n",Arrays.toString(numsL));
        int ans = mergeSortPairCount(numsL,0, nums.length);
        //System.out.printf("Sorted=%s\n",Arrays.toString(numsL));
        return ans;
    }

    public int mergeSortPairCount(long[] numsL,int start, int end){
        //System.out.printf("nums=%s start=%d end=%d\n",Arrays.toString(nums),start,end);
        // base case
        if((end-start)<=1){
            return 0;
        }
        // cmd logic
        int mid = start + (end-start)/2;
        int leftC = mergeSortPairCount(numsL,start,mid);
        int rightC = mergeSortPairCount(numsL,mid,end);
        int currentCount = mergeCount(numsL,start,mid,end);
        return leftC+rightC+currentCount;
    }

    public int mergeCount(long[] numsL, int start, int mid, int end){
        //System.out.printf("merge nums=%s start=%d mid=%d end=%d\n",Arrays.toString(numsL),start,mid,end);
        int l = start;
        int r = mid;

        // Count the reverse pairs in this merge
        int rcount = 0;
        while(l<mid && r<end){
            if(numsL[l]>(2*numsL[r])){
                rcount += (mid-l);
                r++;
            }else{
                l++;
            }
        }
        long[] mergeArr = new long[end-start];
        l = start;
        r = mid;
        int counter=0;
        while(l<mid && r<end){
            if(numsL[l]<=numsL[r]){
                mergeArr[counter] = numsL[l];
                l++;
            }else{
                mergeArr[counter] = numsL[r];
                r++;
            }
            counter++;
        }

        while(l<mid){
            mergeArr[counter] = numsL[l];
            l++;
            counter++;
        }

        while(r<end){
            mergeArr[counter] = numsL[r];
            r++;
            counter++;
        }

        for(int i=0; i<mergeArr.length;i++){
            numsL[start+i] = mergeArr[i];
        }
        //System.out.printf("merge nums=%s start=%d mid=%d end=%d incr=%d\n",Arrays.toString(numsL),start,mid,end,increment);

        return rcount;
    }
}