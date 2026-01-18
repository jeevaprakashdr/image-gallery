import fetchImages from "@/lib/fetchImages"
import { Image } from "@/models/Image"

export default async function Gallery() {

    const images: Image[] | undefined= await fetchImages();

    if(!images)
        return <h2 className="m4 text-2xl font-bold">No Images found</h2>
    
    return (
        <section className="px-2 my-3 grid grid-cols-3 gap-5">
            {images.map(x => (
                <div key={x.id} className="h-64 bg-gray-200 rounded-xl">{x.title}</div>
            ))}
        </section>
    )
}