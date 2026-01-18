import { Image } from "@/models/Image";

export default async function fetchImages(topic: string | undefined)
: Promise<Image[] | undefined> {

    try {
        const url = !topic
            ? "http://localhost:8080/images"
            : `http://localhost:8080/images/search?tag=${topic}`;
        
        const res = await fetch(url)

        if(!res.ok) throw new Error("Failed to fetch image data")
        
        const images : Image[] = await res.json()

        console.log(images)
        
        return images
    } catch(e) {
        if(e instanceof Error)
            console.log(e.stack)
    }
}