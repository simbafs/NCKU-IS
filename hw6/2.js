let table = {};
let max = 11;

for (let a = 1; a < max; a++) {
	let row = {};
	let flag = false;
	for (let n = 1; n < max; n++) {
		let t = Math.pow(a, n) % max;
		row[n] = t
		if (t == 1) {
			if (!flag && n == max - 1) {
				flag = !flag;
				row.isPrime = "✔️";
			}
			flag = true;
		}
	}
	table[a] = row
}

console.table(table);
