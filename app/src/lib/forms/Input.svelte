<script lang="ts">
    import type {HTMLInputTypeAttribute} from "svelte/elements";

    interface Props {
        label?: string
        name: string
        type: HTMLInputTypeAttribute | null | undefined
    }

    let { label, type }: Props = $props()
</script>

<label>
    {label}
    <input class="appearance-none border border-gray-300 outline-none p-2 rounded-sm w-full" {type}>
</label>