<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';

	let { data } = $props();
	let deck = $state(data.deck);
	let page = $state(0);
	let flipedByDefaut = $state(false);
	let fliped = $state(false);

	const flip = () => (fliped = !fliped);
	const next = () => {
		page = (page + 1) % deck.cards.length;
		fliped = flipedByDefaut;
	};
	const prev = () => {
		page = (page + deck.cards.length - 1) % deck.cards.length;
		fliped = false;
	};
	const swapSides = () => {
		flipedByDefaut = !flipedByDefaut;
		fliped = flipedByDefaut;
	};

	const copy = (value: string) => {
		navigator.clipboard.writeText(value).then(() => {
			alert('Copied to clipboard');
		});
	};
</script>

<div class="flex h-full w-full justify-center">
	<div class="flex w-full max-w-[40rem] flex-col gap-4 pt-6">
		<div class="pb-2">
			<h1 class="text-2xl font-medium">
				{deck.title}
			</h1>
			<button
				on:click={() => copy(window.location.href)}
				class="text-muted-foreground font-mono text-sm"
			>
				id:{deck.id}
			</button>
		</div>
		<button
			on:click={flip}
			class:ring-4={fliped}
			class="ring-border bg-background flex aspect-[2/1] w-full shrink-0 items-center justify-center rounded-2xl border p-8 ring-inset transition-transform hover:scale-105 active:scale-100"
		>
			<p class="text-xl">
				{deck.cards[page][Number(fliped)]}
			</p>
		</button>
		<div class="flex items-center">
			<div class="w-full flex-col">
				<p class:opacity-20={fliped} class="font-mono text-sm">front</p>
				<p class:opacity-20={!fliped} class="font-mono text-sm">back</p>
			</div>
			<div class="flex w-full items-center justify-center gap-4">
				<Button variant="outline" size="icon" on:click={prev} class="h-12 w-12 rounded-full">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						width="24"
						height="24"
						viewBox="0 0 24 24"
						fill="none"
						stroke="currentColor"
						stroke-width="2"
						stroke-linecap="round"
						stroke-linejoin="round"
						class="lucide-icon lucide lucide-arrow-left h-4 w-4"
						><path d="m12 19-7-7 7-7"></path><path d="M19 12H5"></path></svg
					>
				</Button>
				<div class="font-mono">{page + 1}/{deck.cards.length}</div>
				<Button variant="outline" size="icon" on:click={next} class="h-12 w-12 rounded-full">
					<svg
						xmlns="http://www.w3.org/2000/svg"
						width="24"
						height="24"
						viewBox="0 0 24 24"
						fill="none"
						stroke="currentColor"
						stroke-width="2"
						stroke-linecap="round"
						stroke-linejoin="round"
						class="lucide-icon lucide lucide-arrow-right h-4 w-4"
						><path d="M5 12h14"></path><path d="m12 5 7 7-7 7"></path></svg
					>
				</Button>
			</div>
			<div class="flex w-full justify-end">
				<Button variant="outline" on:click={swapSides}
					>{flipedByDefaut ? 'back first' : 'front first'}</Button
				>
			</div>
		</div>
		<details open>
			<summary class="select-none pb-2 pt-4 text-lg font-medium hover:cursor-pointer"
				>List of cards</summary
			>
			<ul class="flex flex-col gap-4 pb-12">
				{#each deck.cards as card, i}
					<li class="bg-card flex w-full divide-x overflow-hidden rounded-lg border p-6">
						<p class="w-full p-2 pr-8">{card[0]}</p>
						<p class="w-full p-2 pl-8">{card[1]}</p>
					</li>
				{/each}
			</ul>
		</details>
	</div>
</div>
