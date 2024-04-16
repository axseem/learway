import { apiUrl } from '$lib/api.js';
import type { Deck } from '$lib/types.js';

export const load = async ({ fetch, params }) => {
    const r = await fetch(apiUrl + '/decks/' + params.title, {
        method: 'get'
    });
    const data: Array<Deck> = await r.json();

    return { decks: data };
};