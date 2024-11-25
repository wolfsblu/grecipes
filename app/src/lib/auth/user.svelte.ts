import {fetchProfile as apiFetchProfile, login as apiLogin} from "../api/client";

let profile: User | null = $state(null)

export const createUser = () => {
    const login = async (credentials: Credentials) => {
        const response = await apiLogin(credentials)
        if (response.error) {
            throw response.error
        } else {
            profile = response.data
        }
    }

    const fetchProfile = async () => {
        const response = await apiFetchProfile()
        if (response.error) {
            throw response.error
        } else {
            profile = response.data
        }
    }

    return {
        get profile() {
            return profile
        },
        fetchProfile,
        login,
    }
}