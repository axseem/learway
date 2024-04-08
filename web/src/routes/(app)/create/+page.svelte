<script lang="ts">
	import Input from '$lib/components/ui/input/input.svelte';
	import Button from '$lib/components/ui/button/button.svelte';
	import { apiUrl } from '$lib/api.js';

	type deckCreateParams = {
		title: string;
		cards: [string, string][];
	};

	async function submit(body: deckCreateParams) {
		await fetch(apiUrl + '/deck', {
			method: 'POST',
			credentials: 'include',
			body: JSON.stringify(body)
		}).then((r) => {
			if (r.status >= 400 && r.status < 500) {
				r.json().then((data) => alert(data.message));
			} else if (r.status >= 200 && r.status < 300) {
				r.json().then((data) => window.location.replace('/deck/' + data.id));
			}
		});
	}

	let deck: deckCreateParams = $state({
		title: '',
		cards: [
			['', ''],
			['', '']
		]
	});
</script>

<div class="flex justify-center w-full h-full">
	<div class="flex w-full max-w-[40rem] flex-col gap-4 pt-6">
		<div class="pb-2">
			<h1 class="text-2xl font-medium">Create new deck</h1>
		</div>
		<div class="flex flex-col gap-1.5">
			<label
				class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
				for="title">Title</label
			>
			<Input id="title" name="title" bind:value={deck.title} />
		</div>
		{#each deck.cards as card, i}
			<div class="flex flex-col gap-1.5">
				<p
					class="text-sm font-medium leading-none peer-disabled:cursor-not-allowed peer-disabled:opacity-70"
				>
					#{i + 1}
				</p>
				<li class="flex w-full">
					<textarea
						bind:value={deck.cards[i][0]}
						placeholder="Content for the front..."
						class="w-full p-8 border rounded-l-lg resize-none bg-card"
					/>
					<textarea
						bind:value={deck.cards[i][1]}
						placeholder="Content for the back..."
						class="w-full p-8 border-r rounded-r-lg resize-none bg-card border-y"
					/>
				</li>
			</div>
		{/each}
		<div class="flex place-content-between">
			<Button variant="secondary" on:click={() => deck.cards.push(['', ''])}>Add card</Button>
			<Button on:click={() => submit(deck)}>Create</Button>
		</div>
	</div>
</div>
