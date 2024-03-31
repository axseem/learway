<script>
	import '../app.pcss';
	import Button from '$lib/components/ui/button/button.svelte';
	import Input from '$lib/components/ui/input/input.svelte';

	let { children } = $props();
</script>

<div class="flex h-screen">
	<aside class="flex h-full w-60 shrink-0 flex-col border-r">
		<a class="border-b p-4 font-mono" href="/">learway</a>
		<div class="px-4 pt-4">
			<Input placeholder="Search for knowledge" />
		</div>
		<div class="flex flex-col gap-2 p-4">
			<Button variant="ghost" href="/" class="justify-start">Recent</Button>
			<Button variant="ghost" href="/" class="justify-start">Popular</Button>
			<Button variant="ghost" href="/" class="justify-start">Latest</Button>
		</div>
		<div class="mt-auto flex flex-col p-4">
			<Button href="/create">Create deck</Button>
		</div>
	</aside>
	<div class="bg-muted/40 h-screen w-full scroll-auto p-4">
		{@render children()}
	</div>
</div>
