<script lang="ts">
    import {createRouter} from "./lib/router.svelte";
    import {createUser} from "./lib/auth/user.svelte";
    import {fetchProfile} from "./lib/api/client";

    const router = createRouter()
    let Page = $derived(router.page)

    const user = createUser()
    fetchProfile()
        .then(profile => user.login(profile.data))
        .finally(() => router.registerRoutes()) // Routes need to check for logged-in user so we load them after we fetch the profile
</script>

<Page/>