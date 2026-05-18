sed -i 's|\\/\\/ +kubebuilder:validation:Type=object|\\/\\/ +kubebuilder:validation:XPreserveUnknownFields\\n\\/\\/ +kubebuilder:validation:Type=object|g' apis/vertexai/v1alpha1/generate.sh
