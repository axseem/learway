<script>
	import Input from '$lib/components/ui/input/input.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
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
			<Input placeholder="Email" />
			<Input placeholder="Password" />
		</div>
		<div class="flex flex-col w-full gap-4">
			<Button class="w-full">Log In</Button>
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
		<a class="pt-2 text-sm text-muted-foreground" href="/signup">Create an account</a>
		<a class="pt-2 text-sm text-muted-foreground" href="/recover">Forgot password</a>
	</div>
</div>
