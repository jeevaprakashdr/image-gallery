
import fetchImages from "@/lib/fetchImages"
import { Image } from "@/models/Image"
import ImgContainer from "./ImgContainer";

type Props = {
    topic?: string | undefined
}

export default async function Gallery({ topic } : Props) {
    const images: Image[] | undefined= await fetchImages(topic);
    
    if(!images)
        return <h2 className="m4 text-2xl font-bold">No Images found</h2>
    
    return (
        <section className="px-2 my-3 grid grid-cols-3 gap-5">
            {images.map((x, i) => (
                <ImgContainer key={i} photo={x}/>
            ))}
        </section>
    )
}