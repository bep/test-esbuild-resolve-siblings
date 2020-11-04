import { format } from 'date-fns';

export function hello() {
	let today = format(new Date(), "'Today is a' iiii");
	return `Hello3 from module. Date from date-fns: ${today}`;
}
