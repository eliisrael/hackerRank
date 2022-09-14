//https://www.hackerrank.com/challenges/utopian-tree/problem?isFullScreen=true

const n = 4

function utopianTree(n) {
let height = 1


    for (let i = 1; i <= n; i++) {

        if(i%2 == 0){
            height = height + 1

        }else{
            height = height*2
        }

    }

   return height

}

 console.log(utopianTree(n))