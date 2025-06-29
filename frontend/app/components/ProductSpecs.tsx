import { Card, CardContent } from "@/components/ui/card";

const specs = [{}];

export default function ProductSpecs() {
  return (
    <Card>
      <CardContent className="p-6">
        <h3 className="text-lg font-semibold mb-4">Características técnicas</h3>
        {/* <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          {specs.map((spec, index) => (
            <div
              key={index}
              className="flex justify-between py-2 border-b border-gray-100 last:border-b-0"
            >
              <span className="text-gray-600">{spec.label}</span>
              <span className="font-medium text-right">{spec.value}</span>
            </div>
          ))}
        </div> */}
      </CardContent>
    </Card>
  );
}
