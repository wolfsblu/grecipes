<script lang="ts">
    import type {HTMLInputTypeAttribute} from "svelte/elements";

    interface Props {
        type: HTMLInputTypeAttribute | null | undefined
    }

    let { type }: Props = $props()
</script>

<input {type}>