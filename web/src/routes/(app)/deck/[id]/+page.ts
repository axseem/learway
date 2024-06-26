import { apiUrl } from '$lib/api.js';
import type { Deck } from '$lib/types.js';

export const load = async ({ fetch, params }) => {
    const r = await fetch(apiUrl + '/deck/' + params.id, {
        method: 'get'
    });
    const deck: Deck = await r.json();

    return { deck: deck };
};
