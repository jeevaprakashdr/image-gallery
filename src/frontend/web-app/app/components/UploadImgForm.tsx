'use client'

import { useState, FormEvent, useRef } from "react"

export default function Gallery() {
    const [filePath, setFilePath] = useState('')
    const [title, setTitle] = useState('')
    const [tags, setTags] = useState('')
    
    const fileInputRef = useRef<HTMLInputElement>(null)
    
    const handleTextInputClick = () => {
        fileInputRef.current?.click()
    }

    const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files && e.target.files.length > 0) {
            setFilePath(e.target.files[0].name)
        }
    }

    const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault()

        const formData = new FormData()
        if (fileInputRef.current?.files?.[0]) {
            formData.append('payload', fileInputRef.current.files[0])
        }
        formData.append('title', title)
        formData.append('tags', tags)

        await fetch('http://localhost:8080/images/upload', {
            method: 'POST',
            body: formData,
        })

        setFilePath('')
        setTitle('')
        setTags('')
    }

    return (
        <form onSubmit={handleSubmit} encType="multipart/form-data">
            <input
                type="text"
                value={filePath}
                onClick={handleTextInputClick}
                readOnly
                placeholder="Select File"
                className="bg-blue-100 p-2 text-xl rounded-xl text-black border-0 mr-2"
            />
            <input
                type="file"
                ref={fileInputRef}
                style={{ display: 'none' }}
                onChange={handleFileChange}
            />
            <input
                type="text"
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                placeholder="Title"
                className="bg-blue-100 p-2 text-xl rounded-xl text-black mr-2"
            />
            <input
                type="text"
                value={tags}
                onChange={(e) => setTags(e.target.value)}
                placeholder="Tags"
                className="bg-blue-100 p-2 text-xl rounded-xl text-black mr-2"
            />
            <button className="bg-amber-500 p-2 w-40 text-xl rounded-xl text-black">Upload</button>
        </form>
    )
}