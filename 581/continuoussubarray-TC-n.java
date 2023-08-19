class Solution {
    public int findUnsortedSubarray(int[] nums) {
        Stack<Integer> minIndex = new Stack<Integer>();
        minIndex.push(0);
        boolean flagAdd = true;
        for(int i=1;i<nums.length;i++){
            if(nums[i]<nums[i-1]){
                //System.out.printf("l1: Stack=%s\n",minIndex.toString());
                while(!minIndex.isEmpty() && nums[i]<nums[minIndex.peek()]){
                    minIndex.pop();
                }
                flagAdd = false;
                //System.out.printf("l2: Stack=%s\n",minIndex.toString());
                if(minIndex.isEmpty()){
                    minIndex.push(-1);
                    break;
                }
            }else{
                if(flagAdd){
                    minIndex.push(i);
                }
            }
        }
        //System.out.printf("final l2: Stack=%s\n",minIndex.toString());
        if(flagAdd){
            return 0;
        }

        Stack<Integer> maxIndex = new Stack<Integer>();
        maxIndex.push(nums.length-1);
        boolean flagAddMax = true;

        for(int i= nums.length-2;i>=0;i--){
            if(nums[i]>nums[i+1]){
                //System.out.printf("l3: Stack=%s\n",maxIndex.toString());
                while(!maxIndex.isEmpty() && nums[i]>nums[maxIndex.peek()]){
                    maxIndex.pop();
                }
                flagAddMax = false;
                //System.out.printf("l4: Stack=%s\n",maxIndex.toString());
                if(maxIndex.isEmpty()){
                    maxIndex.push(nums.length);
                    break;
                }
                
            }else{
                if(flagAddMax){
                    maxIndex.push(i);
                }
            }
        }
        //System.out.printf("final l4: Stack=%s\n",maxIndex.toString());
        return (maxIndex.peek()-1)-(minIndex.peek()+1)+1;
    }
}