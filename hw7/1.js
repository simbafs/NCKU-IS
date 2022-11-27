const M = 1121;
const d = 1433;
const n = 3599;

let sign = M % n;

for (let i = 1; i < d; i++) {
	sign = (sign * (M % n)) % n;
}

console.log(sign);
