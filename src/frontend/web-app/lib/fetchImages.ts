import { Image } from "@/models/Image";

export default async function fetchImages()
: Promise<Image[] | undefined> {

    try {
        const res = await fetch("http://localhost:8080/images")

        if(!res.ok) throw new Error("Failed to fetch image data")
        
        const images : Image[] = await res.json()

        console.log(images)
        
        return images
    } catch(e) {
        if(e instanceof Error)
            console.log(e.stack)
    }
}