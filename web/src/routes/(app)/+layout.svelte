<script>
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';

	let { children } = $props();
</script>

<div class="flex h-screen bg-background text-foreground">
	<aside class="flex flex-col h-full border-r w-60 shrink-0">
		<a class="p-4 font-mono border-b" href="/">learway</a>
		<div class="px-4 pt-4">
			<Input placeholder="Search for knowledge" />
		</div>
		<div class="flex flex-col gap-2 p-4">
			<Button variant="ghost" href="/" class="justify-start">Recent</Button>
			<Button variant="ghost" href="/" class="justify-start">Popular</Button>
			<Button variant="ghost" href="/" class="justify-start">Latest</Button>
		</div>
		<div class="flex flex-col gap-2 p-4 mt-auto">
			<Button variant="secondary" href="/login">Log In</Button>
			<Button href="/create">Create deck</Button>
		</div>
	</aside>
	<div class="w-full h-screen p-4 overflow-auto">
		{@render children()}
	</div>
</div>
