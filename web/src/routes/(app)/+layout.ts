import { apiUrl } from '$lib/api.js';
import type { LayoutLoad } from './$types';

export const load: LayoutLoad = async ({ fetch }) => {
	const r = await fetch(apiUrl + '/session', {
		method: 'get',
		credentials: 'include'
	});

	const data = await r.json();

	if (r.status >= 200 && r.status < 300) {
		const expires = new Date(data.expiresAt).toUTCString();
		document.cookie = `sessionID=${data.id}; expires=${expires}; path=/; secure=true;`;
		return { authorized: true };
	} else {
		document.cookie = `sessionID=; max-age=0; path=/; secure=true;`;
		localStorage.removeItem('username');
		return { authorized: false };
	}
};
