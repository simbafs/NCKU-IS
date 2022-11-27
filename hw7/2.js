const p = 19;
const g = 10;
const Yb = 3;

let Xb = 0;

for(let i = 0; ; i++){
	if(Math.pow(g, i)%p === Yb) {
		console.log(`Xb = ${i}`);
		Xb = i;
		break;
	}
}

const k = 5;
const M = 11;

const r = Math.pow(g, k) % p;

for (let s = 0; s < 1e5; s++) {
	if (M === (Xb * r + k * s) % (p - 1)) {
		console.log(`(r, s) = (${r}, ${s})`);
		break;
	}
}
