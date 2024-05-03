package templates

func (t *Templates) HttpClient() string {
	t.setReplacements()

	template := `
	class TangoClient {
		baseUrl: string;
	
		constructor(baseUrl: string) {
			this.baseUrl = baseUrl;
		}
	
		async get(url: string): Promise<any> {
			return this.sendRequest('GET', url);
		}
	
		async put(url: string, data?: any): Promise<any> {
			return this.sendRequest('PUT', url, data);
		}
	
		async delete(url: string): Promise<any> {
			return this.sendRequest('DELETE', url);
		}

		async postFile(url: string, file: File): Promise<any> {
			const formData = new FormData();
			formData.append('file', file);
	
			return this.sendRequest('POST', url, formData);
		}
	
		private async request(method: string, url: string, data?: any): Promise<any> {
		`
	template += "	const response = await fetch(`${this.baseUrl}${url}`, {"
	template += `
				method: method,
				headers: data instanceof FormData ? {} : { 'Content-Type': 'application/json' },
				body: data
			});
			return await response.json();
		}
	
		private async sendRequest(method: string, url: string, data?: any): Promise<any[]> {
			try {
				const response = await this.request(method, url, data);
				return [response, null];
			} catch (error) {
				return [null, error];
			}
		}
	}
	
	// Ejemplo de uso
	/*
	const tangoClient = new TangoClient('https://api.example.com');
	
	tangoClient.get('/users')
		.then(([data, error]) => {
			if (error) {
				console.error('Error en GET:', error);
			} else {
				console.log('GET:', data);
			}
		});
	
	tangoClient.put('/users/123', { name: 'John Doe', age: 30 })
		.then(([data, error]) => {
			if (error) {
				console.error('Error en PUT:', error);
			} else {
				console.log('PUT:', data);
			}
		});
	
	tangoClient.delete('/users/123')
		.then(([data, error]) => {
			if (error) {
				console.error('Error en DELETE:', error);
			} else {
				console.log('DELETE:', data);
			}
		});

	// Envío de archivo
	const fileInput = document.getElementById('fileInput') as HTMLInputElement;
	if (fileInput.files && fileInput.files.length > 0) {
		const file = fileInput.files[0];
		tangoClient.postFile('/upload', file)
			.then(([data, error]) => {
				if (error) {
					console.error('Error en POST de archivo:', error);
				} else {
					console.log('Respuesta de POST de archivo:', data);
				}
		});
	}
	*/
	
	`
	return t.Replacements.Replace(template)
}
