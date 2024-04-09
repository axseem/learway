import { apiUrl } from '$lib/api.js';
import type { Deck } from '$lib/types.js';

export const load = async ({ fetch, params }) => {
    let r = await fetch(apiUrl + '/deck/' + params.id, {
        method: 'get'
    });
    let deck: Deck = await r.json();

    return { deck: deck };
};
