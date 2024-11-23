<script lang="ts">
    import {createRouter} from "./lib/router.svelte";
    import {createUser} from "./lib/auth/user.svelte";
    import {fetchProfile} from "./lib/api/client";

    const router = createRouter()
    router.registerRoutes()
    let Page = $derived(router.page)

    const user = createUser()
    const profile = fetchProfile()
        .then(profile => user.login(profile.data))
</script>

<Page/>