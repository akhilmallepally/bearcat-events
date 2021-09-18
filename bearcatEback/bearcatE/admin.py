from django.contrib import admin
from .models import bearcatE

class bearcatE_admin(admin.ModelAdmin):
    list_display = ('title', 'description', 'completed')
# Register your models here.

admin.site.register(bearcatE, bearcatE_admin)

