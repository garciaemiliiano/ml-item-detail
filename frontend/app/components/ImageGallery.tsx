export default function ImageGallery({ image }: { image: string }) {
  return (
    <div className="rounded-lg overflow-hidden">
      <img src={image} alt="Producto" className="w-full object-cover" />
    </div>
  );
}
