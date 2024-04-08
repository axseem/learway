import { apiUrl } from '$lib/api.js';
import type { Deck } from '$lib/types.js';

export const load = async ({ fetch }) => {
    let r = await fetch(apiUrl + '/decks', {
        method: 'get'
    });
    let data: Array<Deck> = await r.json();

    return { decks: data };
};
