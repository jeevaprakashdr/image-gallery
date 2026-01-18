import fetchImages from "@/lib/fetchImages"
import { Image } from "@/models/Image"
import ImgContainer from "./ImgContainer";

export default async function Gallery() {

    const images: Image[] | undefined= await fetchImages();

    if(!images)
        return <h2 className="m4 text-2xl font-bold">No Images found</h2>
    
    return (
        <section className="px-2 my-3 grid grid-cols-3 gap-5">
            <div className="h-64 bg-gray-200 flex items-center justify-center text-8xl">+</div>
            {images.map((x, i) => (
                <ImgContainer key={i} photo={x}/>
            ))}
        </section>
    )
}