import { apiUrl } from '$lib/api.js';
import type { Deck } from '$lib/types.js';
import type { PageLoad } from "./$types";


export const load: PageLoad = async ({ fetch }) => {
	const r = await fetch(apiUrl + '/decks', {
		method: 'get'
	});
	let decks: Array<Deck> = await r.json();
	decks = decks?.slice(0, 4)

	return { decks: decks };
};
