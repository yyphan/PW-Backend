---
updatedAt: "2025-12-16"
---

Ray tracing is essentially a geometric problem. We shoot a ray from the **camera**, and check if it hits anything in the scene.

> "Mathematics is the music of reason." —— James Joseph Sylvester

## 1. The Math Model

A ray can be defined as $P(t) = O + tD$, where:
- `O` is the origin
- `D` is the direction (normalized)
- `t` is the distance

Here is how we define it in HLSL:

```csharp
struct Ray {
    float3 origin;
    float3 direction;
};
```

Why do we need this? Because frameworks like **HDRP** hide these details
