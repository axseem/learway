<script lang="ts">
	let { data } = $props();
	let user = $state(data.user);
</script>

<p>user @{user.username}</p>
