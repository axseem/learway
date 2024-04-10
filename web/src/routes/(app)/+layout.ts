import type { PageLoad } from "./$types";

export const load: PageLoad = async () => {
    const sessionID = document.cookie
        .split("; ")
        .find((row) => row.startsWith("sessionID="))
        ?.split("=")[1];

    return { authorized: sessionID != undefined }
};