<script lang="ts">
    import {createRecipe, type Recipe} from "../../api/client";
    import Input from "../forms/Input.svelte";

    const recipe: Recipe = $state({
        name: '',
        servings: null,
        minutes: null,
    })

    const saveData = (e: SubmitEvent) => {
        e.preventDefault()
        createRecipe(recipe)
    }
</script>

<form onsubmit={saveData}>
    <Input type="text" label="Name" bind:value={recipe.name}/>
    <Input type="number" label="Servings" bind:value={recipe.servings}/>
    <Input type="number" label="Minutes" bind:value={recipe.minutes}/>

    <button type="submit">
        Create
    </button>
</form>