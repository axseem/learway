<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';

	let { data } = $props();
	let deck = $state(data.deck);
	let page = $state(0);
	let fliped = $state(false);

	const flip = () => (fliped = !fliped);
	const next = () => {
		page = (page + 1) % deck.cards.length;
		fliped = false;
	};
	const prev = () => {
		page = (page + deck.cards.length - 1) % deck.cards.length;
		fliped = false;
	};

	let cardEl: HTMLButtonElement;
</script>

<div class="flex h-full w-full justify-center">
	<div class="flex w-full max-w-[40rem] flex-col gap-4 pt-6">
		<h1 class="pb-2 text-3xl font-medium first:mt-0">
			{deck.title}
		</h1>
		<button
			bind:this={cardEl}
			on:click={flip}
			class:ring-4={fliped}
			class="ring-border flex aspect-[2/1] w-full items-center justify-center rounded-2xl border bg-white p-8 ring-inset transition-transform hover:scale-105 active:scale-100"
		>
			<p>
				{deck.cards[page][Number(fliped)]}
			</p>
		</button>
		<div class="flex items-center justify-center gap-4">
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
	</div>
</div>
