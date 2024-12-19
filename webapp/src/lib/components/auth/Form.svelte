<script lang="ts">
    import type {Component, Snippet} from "svelte";
    import Button from "../forms/Button.svelte";

    interface Props {
        children: Snippet
        imageSrc: string
        submitIcon: Component
        submitLabel: string
        onSubmit: (e: SubmitEvent) => void
        subtitle: string
        title: string
    }

    let {children, imageSrc, title, submitIcon, submitLabel, subtitle, onSubmit}: Props = $props()
</script>

<div class="bg-white lg:grid grid-cols-[1fr_2fr] md:shadow-lg md:w-1/3">
    <div class="bg-center bg-cover rounded-l-md"
         style="background-image: url({imageSrc})">
    </div>
    <form class="md:border border-gray-200 grid grid-rows-[auto_min-content] rounded-r-md"
          onsubmit="{onSubmit}">
        <div class="p-6">
            <h1 class="font-light mb-3 text-3xl">{title}</h1>
            <p class="mb-3 text-sm">
                {@html subtitle}
            </p>
            {@render children()}
        </div>

        <div class="p-6 pt-0">
            <Button type="submit" class="mt-3" icon={submitIcon}>
                {submitLabel}
            </Button>
        </div>
    </form>
</div>