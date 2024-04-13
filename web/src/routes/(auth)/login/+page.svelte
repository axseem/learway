<script lang="ts">
	import Input from '$lib/components/ui/input/input.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import type { UserCredentials } from '$lib/types';
	import { apiUrl } from '$lib/api';

	const login = async (p: UserCredentials) => {
		let r = await fetch(apiUrl + '/login', {
			method: 'post',
			body: JSON.stringify(p),
			headers: {
				'Content-Type': 'application/json'
			}
		});

		if (r.status >= 400 && r.status < 500) {
			let data = await r.json();
			alert(data.message);
			return;
		}

		if (r.status >= 200 && r.status < 300) {
			let data = await r.json();
			let session = data.session;
			let username = data.username;
			let expires = new Date(session.expiresAt).toUTCString();
			document.cookie = `sessionID=${session.id}; expires=${expires}; path=/; secure=true;`;
			localStorage.setItem('username', username);
			window.location.replace('/');
			return;
		}
	};

	let credantials: UserCredentials = $state({
		email: '',
		password: ''
	});
</script>

<div class="flex w-full flex-col gap-2">
	<form
		class="bg-card text-card-foreground flex w-full flex-col gap-6 rounded-lg border p-6 shadow-sm"
		on:submit={() => {
			login(credantials);
		}}
	>
		<div class="flex flex-col gap-1.5">
			<h3 class="text-2xl font-semibold tracking-tight">Welcome</h3>
			<p class="text-muted-foreground text-sm">Enter your credantials below</p>
		</div>
		<div class="flex w-full flex-col gap-2">
			<Input bind:value={credantials.email} placeholder="Email" />
			<Input bind:value={credantials.password} placeholder="Password" />
		</div>
		<div class="flex w-full flex-col gap-4">
			<Button type="submit" class="w-full">Log in</Button>
			<div class="relative">
				<div class="absolute inset-0 flex items-center"><span class="w-full border-t"></span></div>
				<div class="relative flex justify-center text-xs uppercase">
					<span class="bg-background text-muted-foreground px-2">Or continue with</span>
				</div>
			</div>
			<Button variant="outline" class="w-full">Continue with Google</Button>
		</div>
	</form>
	<div class="flex place-content-between">
		<a class="text-muted-foreground pt-2 text-sm" href="/signup">Create an account</a>
		<a class="text-muted-foreground pt-2 text-sm" href="/recover">Forgot password</a>
	</div>
</div>
