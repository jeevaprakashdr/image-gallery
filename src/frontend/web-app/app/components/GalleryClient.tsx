"use client"

import { useEffect, useState } from "react";
import { Image } from "@/models/Image";
import fetchImages from "@/lib/fetchImages";
import ImgContainer from "./ImgContainer";

type Props = {
    topic?: string | undefined
}

export default function GalleryClient({ topic }: Props) {
    const [images, setImages] = useState<Image[] | undefined>(undefined);

    useEffect(() => {
        fetchImages(topic).then(setImages);

        const ws = new WebSocket("ws://localhost:8081/ws");

        ws.onmessage = (event) => {
            const newImage = new Image();
            newImage.id = event.data
            setImages(prev =>
                prev ? [...prev, newImage] : [newImage]
            );
        };

        return () => ws.close();
    }, [topic]);

    if (!images)
        return <h2 className="m4 text-2xl font-bold">No Images found</h2>;

    return (
        <section className="px-2 my-3 grid grid-cols-3 gap-5">
            {images.map((x, i) => (
                <ImgContainer key={i} photo={x} />
            ))}
        </section>
    );
}