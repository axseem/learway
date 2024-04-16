import { apiUrl } from '$lib/api.js';

export const load = async ({ fetch, params }) => {
    const r = await fetch(apiUrl + '/user/' + params.id, {
        method: 'get'
    });
    const user = await r.json();

    return { user: user };
};
