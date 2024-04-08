<script lang="ts">
	import Input from '$lib/components/ui/input/input.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import type { UserCreateParams } from '$lib/types';
	import { apiUrl } from '$lib/api';

	const login = async (p: UserCreateParams) => {
		let r = await fetch(apiUrl + '/signup', {
			method: 'post',
			body: JSON.stringify(p)
		});

		console.log(r);

		if (r.status >= 400 && r.status < 500) {
			r.json().then((data) => alert(data.message));
			return;
		}

		if (r.status >= 200 && r.status < 300) {
			r.json().then((data) => {
				document.cookie = `sessionID=${data.id};expires=${data.expiresAt};path=/`;
			});
			alert('UWEEEEE');
			return;
		}
	};

	let credantials: UserCreateParams = $state({
		username: '',
		email: '',
		password: ''
	});
</script>

<div class="flex flex-col w-full gap-2">
	<form
		class="flex flex-col w-full gap-6 p-6 border rounded-lg shadow-sm bg-card text-card-foreground"
	>
		<div class="flex flex-col gap-1.5">
			<h3 class="text-2xl font-semibold tracking-tight">Welcome</h3>
			<p class="text-sm text-muted-foreground">Enter your credantials below</p>
		</div>
		<div class="flex flex-col w-full gap-2">
			<Input bind:value={credantials.username} placeholder="Username" />
			<Input bind:value={credantials.email} placeholder="Email" />
			<Input bind:value={credantials.password} placeholder="Password" />
		</div>
		<div class="flex flex-col w-full gap-4">
			<Button
				on:click={() => {
					login(credantials);
				}}
				class="w-full">Sign Up</Button
			>
			<div class="relative">
				<div class="absolute inset-0 flex items-center"><span class="w-full border-t"></span></div>
				<div class="relative flex justify-center text-xs uppercase">
					<span class="px-2 bg-background text-muted-foreground">Or continue with</span>
				</div>
			</div>
			<Button variant="outline" class="w-full">Continue with Google</Button>
		</div>
	</form>
	<div class="flex place-content-between">
		<a class="pt-2 text-sm text-muted-foreground" href="/login">Log in</a>
		<a class="pt-2 text-sm text-muted-foreground" href="/recover">Forgot password</a>
	</div>
</div>
