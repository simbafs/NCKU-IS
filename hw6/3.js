let a = 2 
let b = 17
let n = 29

for(let i = 0; ; i++){
	if(Math.pow(a, i)%n === b){
		console.log(i)
		break
	}
}
