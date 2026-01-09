​
# Series Introduction
I am not sure if there is a common name for this effect:

![Light Band Effect](/images/2026/01/20260107-LightBandEffect.gif?h=200)

But I will call it "light band" for now. It is essentially a light band moving across the target surface, and is commonly used to attract user's attention or to demonstrate material properties such as being shiny or reflective.

This series of posts will introduce a few ways to implement this effect in Unity using Shader Graph, and it gets more hacky (and more fun!) as it goes:

1. Using a light band texture, move it across the surface of the object (UV space)

2. Move the light band across the screen, masking it such that it is only visible on the target object (Screen space)

3. Use a program-generated light band in 2

&nbsp;

This post is the first entry in the series. I will go through the 1st method - moving texture in UV space.


# Method 1 (Moving Texture in UV Space)
If you have a light band texture ready with you, the easiest way to implement is to sample it and apply it to the target object. By offsetting and rotating its UV, we can make the light band move however we want.

![Shine Line Tex](/images/2026/01/20260107-ShineLineTex.png?h=200)

## Sampling
In our Shader Graph, we will first sample the base color of the target object. I used a simple `Sample Texture 2D` node, together with a optional `MainColor` variable:

![Base Color](/images/2026/01/20260107-BaseColor.png)

Next, we will sample the light band texture in pretty much the same way. I added an `Alpha` variable to control its transparency, and also added a `Tiling and Offset` node for later use:

![Light Band Color](/images/2026/01/20260107-LightBandColor.png)

## Blending
Having both colors, we can preview the effect. I added these 2 colors (you can experiement with other blending modes) to get a static look of the light band effect.

![Blending](/images/2026/01/20260107-Blend.png?h=600)

## Moving it
Remember the `Tiling and Offset` node we added just now? By adjusting offset, we can move the light band. I defined a `Step` variable to represent the "distance" the light band has traveled on the target object (0 to 1):

![Step](/images/2026/01/20260107-Step.png)

> With a float variable selected, you can select Mode: Slider in the variable window to enable sliding control.

When `Step` changes, we can see the light band moving.

## Rotating it
Most of the time we will also need to set the angle of the light band. Remember we just offsetted the UV? Turns out we can rotate it too. Here I defined an `Angle` variable, and connected it to the UV's `Rotate` node:

![Angle](/images/2026/01/20260107-Angle.png)

The rotated UV is then connected to the `Tiling and Offset` node from the previous step, allowing us to control the rotation of the light band.

# Conclusion
The complete Shader Graph looks like this:

![Full Graph](/images/2026/01/20260107-Full-Graph.png)

Simple, isn't it? But good enough most of the time.

However, due to the fact that everything here happens in UV space, it does have some limitations: 

- UV's not suitable? 
- Don't want to render the light band on the back of the target object?

The rest of the posts in this series aim to address these limitations.
